import requests

# Define the API endpoint and query parameters
def getFile(fileName :str, port :str):
    response = requests.get(f"http://localhost:{port}/getFile?fileName={fileName}")
    if response.status_code == 200:
        print(response.text, end="")


url = "http://localhost:8000/getFile"
params = {"fileName": "server.go"}

try:
    # Make a GET request to the API
    response = requests.get(url, params=params)

    # Check if the request was successful
    if response.status_code == 200:
        # Parse the JSON response
        data = response.json()

        # Iterate over the response and print chunk names and ports
        for entry in data:
            chunk_name = entry.get("ChunkName")
            chunk_server_addr = entry.get("ChunkServerAddr")
            
            # Assuming we take the first port from the ChunkServerAddr list
            if chunk_server_addr:
                port = chunk_server_addr[0].split(":")[-1]
                getFile(chunk_name,port)
            else:
                print(f"ChunkName: {chunk_name}, No ChunkServerAddr found")

    else:
        # Print an error message if the request was not successful
        print(f"Failed to get the file. Status code: {response.status_code}")

except requests.exceptions.RequestException as e:
    print(f"Error while")



















