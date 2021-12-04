# Go Voting App

## Overview
A simple cats vs dogs voting web app written in Go programming langulage and backed by Redis.

## Running on docker
    docker-compose up -d

### POST a vote
    POST /vote/dogs
    # or 
    POST /vote/cats

### GET votes
    GET /votes
