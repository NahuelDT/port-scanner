# Port-Scanner

Port-Scanner is a tool to scan ports of any given host. Complies with all but the fourth iteration.

#### Iteration 1:
 - Scans all ports and reports the ones open.

#### Iteration 2:
 - Uses concurrency to scan the ports, you can customize the amount of threads to use, the default is 5.

#### Iteration 3: 

Avoids getting detected by the firewall:
 - Scans odd and even ports separately so it never scans consecutive ports (Waits 2 seconds between even and odds).
- Uses randomized timeout to avoid opening and closing at least 4 ports in a predictable timely manner.

#### Iteration 5:
 - Dockerized  
 



## Installation
### Using Go
Go run command should install all necessary dependencies.
 
### Using Docker
Build Docker image
```bash
Docker build -t port-scanner .
```

## Usage
### Arguments

| Flags | Description |Default |
| ------ | ------ |------ |
| -host | HostName you want to scan. | google.com|
| -t | Number of threads to use.  | 5 |
| -start | Start of port range to scan.| 0 |
| -end | End of port range to scan.| 1024 |
| -f | If true, is not wary of a firewall and may be detected.| False |


### Go run 
Run main.go using the command scan

```bash
go run main.go scan -host 
```
### Docker Build 
Run the Docker image
```bash
docker run port-scanner:latest -host stackoverflow.com
```

### Examples
#### Using a custom amount of threads
```bash
docker run port-scanner:latest -host stackoverflow.com -t 10 
```
or 

```bash
go run main.go scan -host stackoverflow.com -t 10
```

#### Using a custom range of ports
```bash
docker run port-scanner:latest -host stackoverflow.com -start 60  -end 10000
```
or 

```bash
go run main.go scan -host stackoverflow.com -start 60  -end 10000
```

#### Without usign the Firewall detect prevention.
```bash
docker run port-scanner:latest -host stackoverflow.com -f
```
or 

```bash
go run main.go scan -host stackoverflow.com -f
```



## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
[MIT](https://choosealicense.com/licenses/mit/)