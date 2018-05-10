## Experimental GRPC Server using Go, Docker and Dep (Feel Free to Use this Code)

## Recipes
   - Golang version 1.7+ https://golang.org/
   
   - Dep https://golang.github.io/dep/

      What the heck is Dep
      The Documentation says `the "official experiment" dependency management tool for the Go language`.

   - Protocol Buffer (Protobuf)
      Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler.
      https://github.com/google/protobuf/releases

   - Docker

## Running
  - Using Go
    - Install Dependencies
      ```shell
      $ dep ensure
      ```

    - Build Binary
      - OSX
        - Build
          ```shell
          $ make grpc-awesome-osx
          ```
        - Run
          ```shell
          $ GRPC_AUTH_KEY=123456 ./grpc-awesome-osx
          ```
      - Linux
        - Build
          ```shell
          $ make grpc-awesome-osx
          ```
        - Run
          ```shell
          $ GRPC_AUTH_KEY=123456 ./grpc-awesome-osx
          ```
  - Using Docker
    - Build
      ```shell
      $ make docker
      ```
    - Run
      ```shell
      $ docker run --rm -p 3000:3000 -e GRPC_AUTH_KEY=123456 go-grpc-awesome
      ```

##

2018 Bhinneka.com
