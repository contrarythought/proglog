$url = "localhost:8080"
$jsonData = @{offset = 3} | ConvertTo-Json
$queryParams = @{
    "data" = $jsonData
}
$response = Invoke-RestMethod -Method Get -Uri $url -Query $queryParams