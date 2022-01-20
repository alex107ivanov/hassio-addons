package main

import (
    "fmt"
    mqtt "github.com/eclipse/paho.mqtt.golang"
    "log"
    "os"
    "encoding/json"
    "syscall"
    "os/signal"
    "github.com/alecthomas/jsonschema"
    "github.com/xeipuuv/gojsonschema"
)

type ServerConfig struct {
  Host string `json:"host" jsonschema:"required, type=string"`
  Port int32 `json:"port" jsonschema:"required, type=integer"`
  Username string `json:"username" jsonschema:"required, type=string"`
  Password string `json:"password" jsonschema:"required, type=string"`
  ClientId string `json:"client_id" jsonschema:"required, type=string"`
}

type SrcConfig struct {
  Server ServerConfig `json:"server" jsonschema:"required"`
  Subscribe []string `json:"subscribe" jsonschema:"required"`
}

type DstConfig struct {
  Server ServerConfig `json:"server" jsonschema:"required"`
  Prefix string `json:"prefix" jsonschema:"required, type=string"`
}

type Configuration struct {
  Src SrcConfig `json:"src" jsonschema:"required"`
  Dst DstConfig `json:"dst" jsonschema:"required"`
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    log.Printf("Connected")
}

func main() {
  configFile = "/data/options.json"
  file, _ := os.Open(configFile)
  defer file.Close()
  decoder := json.NewDecoder(file)
  config := Configuration{}
  sc := jsonschema.Reflect(&config)
  b, _ := json.Marshal(sc)
  log.Printf("Schema: %s", string(b))
  err := decoder.Decode(&config)
  if err != nil {
    log.Printf("error: %v", err)
    os.Exit(1)
  }
  log.Printf("Config: %v", config)
  // Check config
  schemaLoader := gojsonschema.NewStringLoader(string(b))
  documentLoader := gojsonschema.NewReferenceLoader("file://" + configFile)
  result, err := gojsonschema.Validate(schemaLoader, documentLoader)
  if err != nil {
    log.Printf("Error loading config: %s", err.Error())
    os.Exit(1)
  }
  if result.Valid() {
    log.Printf("Config is valid")
  } else {
    log.Printf("Config is not valid. see errors:")
    for _, desc := range result.Errors() {
      log.Printf("- %s\n", desc)
    }
    os.Exit(1)
  }

  quit := make(chan os.Signal)

    mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
    mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
    mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
//    mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)



  dstClient, err := createClient(config.Dst.Server, func(client mqtt.Client, msg mqtt.Message){}, func(client mqtt.Client, err error) {
    log.Printf("Connect lost: %v", err);
    select {
      case quit <- os.Interrupt: {}
    }
  })
  if err != nil {
    log.Printf("Error connecting to dst MQTT server: %v", err)
    os.Exit(2)
  }

  srcClient, err := createClient(config.Src.Server, func(client mqtt.Client, msg mqtt.Message){
    log.Printf("Got message %s from topic %s", msg.Payload(), msg.Topic())
    newTopic := config.Dst.Prefix + "/" + msg.Topic()
    token := dstClient.Publish(newTopic, msg.Qos(), msg.Retained(), msg.Payload())
    token.Wait()
    if token.Error() != nil {
      log.Printf("Error sending message to topic %s: %s", newTopic, token.Error())
      select {
        case quit <- os.Interrupt: {}
      }
    }
    log.Printf("Message sent to topic %s", newTopic)
  }, func(client mqtt.Client, err error) {
    log.Printf("Connect lost: %v", err);
    select {
      case quit <- os.Interrupt: {}
    }
  })
  if err != nil {
    log.Printf("Error connecting to src MQTT server: %v", err)
    os.Exit(2)
  }
  err = sub(srcClient, config.Src.Subscribe)
  if err != nil {
    log.Printf("Error subscribing to topics: %s", err)
    os.Exit(3)
  }

  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit

  log.Printf("Exiting")

  if srcClient.IsConnected() {
    srcClient.Disconnect(250)
  }
  if dstClient.IsConnected() {
    dstClient.Disconnect(250)
  }

  log.Printf("Done");
}

func createClient(config ServerConfig, msgHandler mqtt.MessageHandler, connectLostHandler mqtt.ConnectionLostHandler) (mqtt.Client, error) {
    var broker = config.Host
    var port = config.Port
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
    opts.SetClientID(config.ClientId)
    opts.SetUsername(config.Username)
    opts.SetPassword(config.Password)
    opts.SetDefaultPublishHandler(msgHandler)
    opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler
    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        return nil, token.Error()
    }
    return client, nil
}

func sub(client mqtt.Client, subscribe []string) error {
  if len(subscribe) == 0 {
    subscribe = []string{"#"}
  }
  for _, topic := range subscribe {
    token := client.Subscribe(topic, 1, nil)
    token.Wait()
    if token.Error() != nil {
      log.Printf("Error subscribing to topic %s: %s", topic, token.Error())
      return token.Error()
    }
    log.Printf("Subscribed to topic: %s", topic)
  }
  return nil
}
