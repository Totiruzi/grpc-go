#!/bin/bash

protoc --go_out=. --go_opt=paths=source_relative greet/greetpb/greet.proto