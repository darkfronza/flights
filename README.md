# flights
A microservice that sorts a list of flights then outputs the source and final destination path among the list.

## Building

### Linux platforms

On the terminal cd to the project folder then run `make`, an executable named `app` should be generated at the project folder.

Example (build and run):
```
$ cd project_folder
$ make
$ ./app
```


### Other platforms (Linux included)

[Install](https://go.dev/doc/install) the go toolchain, please follow the instructions in this [link](https://go.dev/doc/install).

With go properly installed, open a terminal and run the following commands to build the app:

```
cd path_to_project_folder
go build -o app
```

## API
The flights microservice expects json-encoded *POST* requests on the `/calculate` endpoint.

The body must be in the following format:
```
{
    "flights": [["SRC1", "DST1], ["SRC2", "DST2"], ...]
}
```

Example request in *curl*:
```
$ curl -X POST http://localhost:8080/calculate \
   -H 'Content-Type: application/json' \
   -d '{"flights":[["ATL", "EWR"], ["SFO", "ATL"]]}'

["SFO","EWR"]
```

## Running

On a terminal, run the `app` command built during the building stage.

By default, the http server listens on port 8080.

If another port is required, just set the API_PORT enviroment variable to the desired port number.

Example running on a custom port on Linux: 
```
$ API_PORT=3000 ./app
```

## Running tests

On a terminal, run the `make test` command.

## Algorithm description

The presented solution assumes that the input consists of a list of source/destination pairs, whose data set compose a directed graph.

Here is the solution outline:

1. Create a tracker object, a map that tracks each airport by its key, e.g. "SFO", "ATL", etc.
2. Each entry in the tracker map has a corresponding integer value, whose initial value is zero.
3. Iterate over each source/destination airport pair in the input list.
4. For the source element in the pair, decrement the counter value by 1 in the tracker.
5. For the destination element in the pair, increment the counter value by 1 in the tracker.

For a non-ciclical graph, after processing each element, the tracker map must have one element whose counter is -1, this is the starting airport.

Each element with a value of zero are airports in the middle.

The ending airport will have a value of 1.

By processing the tracker map and looking for the min,max counters we have the starting/endpoint airports.
