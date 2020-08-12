----------------------------------------------

REST API Prototype

----------------------------------------------


# Prerequisites running simprints

Install Golang. Using [GVM](https://github.com/moovweb/gvm) and Golang 1.12.4 is recommended

# Running

```bash
GCLOUD_PROJECT=your-gcp-project go run cmd/mywebapp/main.go
```

# Testing

```bash
GCLOUD_PROJECT=your-gcp-project go test ./... -count=1
```

Need to set the auth credential path 
export GOOGLE_APPLICATION_CREDENTIALS ="/home/user/Downloads/[FILE_NAME].json"
