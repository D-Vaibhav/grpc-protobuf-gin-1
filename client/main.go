package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vaibhav/am/protos"
	"google.golang.org/grpc"
)

func main() {
	// dial to the server address : localhost:4040
	clientConnection, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := protos.NewAddMultiplyServiceClient(clientConnection)

	// creating gin-gonic router for routing purpose
	ginRouter := gin.Default()

	// api end-points
	ginRouter.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, b := ctx.Param("a"), ctx.Param("b")
		toBase := 10
		bitSize := 64

		numberA, err := strconv.ParseUint(a, toBase, bitSize)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		numberB, err := strconv.ParseUint(b, toBase, bitSize)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		// after getting both the numberA, numberB
		req := &protos.Request{A: int64(numberA), B: int64(numberB)}

		if response, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	})

	ginRouter.GET("/multiply/:a/:b", func(ctx *gin.Context) {
		a, b := ctx.Param("a"), ctx.Param("b")
		toBase := 10
		bitSize := 64

		numberA, err := strconv.ParseUint(a, toBase, bitSize)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		numberB, err := strconv.ParseUint(b, toBase, bitSize)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		// after getting both the numberA, numberB
		req := &protos.Request{A: int64(numberA), B: int64(numberB)}

		if response, err := client.Multiply(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

	})

	if err := ginRouter.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
