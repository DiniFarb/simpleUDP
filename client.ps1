$serverAddress = "localhost"
$serverPort = 8866
$message = "Hello, UDP server!"

$udpClient = New-Object System.Net.Sockets.UdpClient
$udpClient.Connect($serverAddress, $serverPort)
$udpData = [System.Text.Encoding]::ASCII.GetBytes($message)
$udpClient.Send($udpData, $udpData.Length)
$udpClient.Close()
