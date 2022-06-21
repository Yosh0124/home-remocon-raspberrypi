package main

import (
	"net/http"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio"
)

const (
	clientID          string = "home-remocon"  // MQTT Client ID
	topicRemoconLight string = "remocon/light" // Topic for controling room light
	topicRemoconAir   string = "remocon/air"   // Topic fpr controling room air conditioner
	pinNumLightOn     int    = 4               // GPIO for turing on a light
	pinNumLightOff    int    = 17              // GPIO for turing off a light
	pinNumLightSmall  int    = 27              // GPIO for turing on a small light
	pinNumAirOff      int    = 18              // GPIO for turing off a air conditioner
	pinNumAirCooler   int    = 5               // GPIO for turing on a cooler
	pinNumAirHeader   int    = 6               // GPIO for turing on a cooler

)

func main() {
	// Initialize godotenv
	if err := godotenv.Load(); err != nil {
		log.Panic(err.Error())
	}

	// Initialize GPIOs
	if err := rpio.Open(); err != nil {
		log.Panic(err.Error())
	}
	defer rpio.Close()

	pinLightOn := rpio.Pin(pinNumLightOn)
	pinLightOn.Output()
	pinLightOn.High()
	pinLightOff := rpio.Pin(pinNumLightOff)
	pinLightOff.Output()
	pinLightOff.High()
	pinLightSmall := rpio.Pin(pinNumLightSmall)
	pinLightSmall.Output()
	pinLightSmall.High()
	pinAirOff := rpio.Pin(pinNumAirOff)
	pinAirOff.Output()
	pinAirOff.High()
	pinAirCooler := rpio.Pin(pinNumAirCooler)
	pinAirCooler.Output()
	pinAirCooler.High()
	pinAirHeater := rpio.Pin(pinNumAirHeader)
	pinAirHeater.Output()
	pinAirHeater.High()

	// //define a function for the default message handler
	// var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	// 	log.Info("TOPIC: %s\n", msg.Topic())
	// 	log.Info("MSG: %s\n", msg.Payload())
	// }

	// // define a function for the room light message handler
	// var hRemoconLight MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	// 	log.Info("TOPIC: %s\n", msg.Topic())
	// 	log.Info("MSG: %s\n", msg.Payload())

	// 	// Unmarshal json
	// 	msgMap := make(map[string]interface{})
	// 	if err := json.Unmarshal(msg.Payload(), &msgMap); err != nil {
	// 		log.Error(err.Error())
	// 	}

	// 	// json format :
	// 	// {
	// 	//   "status" : "on", "off" or "small"
	// 	// }
	// 	switch msgMap["status"] {
	// 	case "on":
	// 		pinLightOn.Low()
	// 		time.Sleep(100 * time.Millisecond)
	// 		pinLightOn.High()
	// 	case "off":
	// 		pinLightOff.Low()
	// 		time.Sleep(100 * time.Millisecond)
	// 		pinLightOff.High()
	// 	case "small":
	// 		pinLightSmall.Low()
	// 		time.Sleep(100 * time.Millisecond)
	// 		pinLightSmall.High()
	// 	}
	// }

	// // define a function for the room light message handler
	// var hRemoconAir MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	// 	log.Info("TOPIC: %s\n", msg.Topic())
	// 	log.Info("MSG: %s\n", msg.Payload())

	// 	// Unmarshal json
	// 	msgMap := make(map[string]interface{})
	// 	if err := json.Unmarshal(msg.Payload(), &msgMap); err != nil {
	// 		log.Error(err.Error())
	// 	}

	// 	// json format :
	// 	// {
	// 	//   "status" : "on", "off" or "small"
	// 	// }
	// 	switch msgMap["status"] {
	// 	case "off":
	// 		pinAirOff.Low()
	// 		time.Sleep(100 * time.Millisecond)
	// 		pinAirOff.High()
	// 	case "cooler":
	// 		pinAirCooler.Low()
	// 		time.Sleep(100 * time.Millisecond)
	// 		pinAirCooler.High()
	// 	case "heater":
	// 		pinAirHeater.Low()
	// 		time.Sleep(100 * time.Millisecond)
	// 		pinAirHeater.High()
	// 	}
	// }

	// // Inisialize MQTT
	// //create a ClientOptions struct setting the broker address, clientid, turn
	// //off trace output and set the default message handler
	// mqttUsername := os.Getenv("MQTT_USER")
	// mqttPassword := os.Getenv("MQTT_PASSWD")
	// mqttHost := os.Getenv("MQTT_HOST")
	// opts := MQTT.NewClientOptions().AddBroker(mqttHost)
	// opts.SetClientID(clientID)
	// opts.SetDefaultPublishHandler(f)
	// opts.SetUsername(mqttUsername)
	// opts.SetPassword(mqttPassword)

	// //create and start a client using the above ClientOptions
	// c := MQTT.NewClient(opts)
	// if token := c.Connect(); token.Wait() && token.Error() != nil {
	// 	log.Panic(token.Error())
	// }
	// defer c.Disconnect(250)

	// //subscribe to the light remocon topic and request messages to be delivered
	// //at a maximum qos of zero, wait for the receipt to confirm the subscription
	// if token := c.Subscribe(topicRemoconLight, 0, hRemoconLight); token.Wait() && token.Error() != nil {
	// 	log.Panic(token.Error())
	// }
	// defer func() {
	// 	//unsubscribe from /go-mqtt/sample
	// 	if token := c.Unsubscribe(topicRemoconLight); token.Wait() && token.Error() != nil {
	// 		log.Panic(token.Error())
	// 	}
	// }()
	// if token := c.Subscribe(topicRemoconAir, 0, hRemoconAir); token.Wait() && token.Error() != nil {
	// 	log.Panic(token.Error())
	// }
	// defer func() {
	// 	//unsubscribe from /go-mqtt/sample
	// 	if token := c.Unsubscribe(topicRemoconAir); token.Wait() && token.Error() != nil {
	// 		log.Panic(token.Error())
	// 	}
	// }()

	// Nextjs
	go func() {
		startFront := exec.Command("yarn", "start")
		startFront.Dir = "front"
		_, err := startFront.Output()
		if err != nil {
			log.Panic(err)
		}
	}()

	// Go
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/light", func(c echo.Context) error {
		// Unmarshal json
		msgMap := make(map[string]interface{})
		if err := c.Bind(&msgMap); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, nil)
		}

		// json format :
		// {
		//   "status" : "on", "off" or "small"
		// }
		switch msgMap["status"] {
		case "on":
			pinLightOn.Low()
			time.Sleep(100 * time.Millisecond)
			pinLightOn.High()
			log.Info("Turn on light.")
		case "off":
			pinLightOff.Low()
			time.Sleep(100 * time.Millisecond)
			pinLightOff.High()
			log.Info("Turn off light.")
		case "small":
			pinLightSmall.Low()
			time.Sleep(100 * time.Millisecond)
			pinLightSmall.High()
			log.Info("Turn small light.")
		}

		return c.JSON(http.StatusOK, msgMap)
	})
	e.POST("/air", func(c echo.Context) error {

		// Unmarshal json
		msgMap := make(map[string]interface{})
		if err := c.Bind(&msgMap); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, nil)
		}

		// json format :
		// {
		//   "status" : "off", "cooler" or "heater"
		// }
		switch msgMap["status"] {
		case "off":
			pinAirOff.Low()
			time.Sleep(100 * time.Millisecond)
			pinAirOff.High()
			log.Info("Turn off air conditioner.")
		case "cooler":
			pinAirCooler.Low()
			time.Sleep(100 * time.Millisecond)
			pinAirCooler.High()
			log.Info("Turn on cooler.")
		case "heater":
			pinAirHeater.Low()
			time.Sleep(100 * time.Millisecond)
			pinAirHeater.High()
			log.Info("Turn on heater.")
		}

		return c.JSON(http.StatusOK, msgMap)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
