# Cleandemo

## Description

Simple TV episode lookup demo application using Go lang to produce a clean web architecture. In the spirit of Go, I have avoided using any frameworks to provide web functionality. Also, to keep things simple a in-memory repository was provided rather than introducing a database dependency. The key tenant of this architecture is that all dependencies must point inward only (towards the domain layer). For more information on clean architecuture along with a better explanation of the aforementioned statement visit the link [here](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html).

## Usage
1. Get all dependencies via 

```
go get -d ./...
```
2. Run the server via
```
go run main.go
```
3. Open browser/PostMan/etc. and call the endpoint with an episode name. 

*Note: X-Files Episode "Kitten" (S11E06) is the only episode currently in the in-memory repository dictionary. All other episode names will return null.*
```
http://localhost:8888/episodes/{episodeName}
```