Import-Module BitsTransfer

$processorArchitecture = (wmic os get osarchitecture)[2].Trim()

$mongoDownloadUrl = "https://fastdl.mongodb.org/win32/mongodb-win32-x86_64-2008plus-ssl-3.4.7-signed.msi"
$mongoInstallPath = "C:\Program Files\MongoDB\Server\3.4"
$mongoInstallFilename = "mongo-install.msi"
$mongoDataPath = "\data\db"

$scriptDir = Split-Path -Path $MyInvocation.MyCommand.Definition -Parent

# set correct download url for processor
if($processorArchitecture -eq "32-bit") {
    Write-Host "Found 32bit processor, downloading Mongo 3.2 x86"
    $mongoDownloadUrl = "http://downloads.mongodb.org/win32/mongodb-win32-i386-v3.2-latest-signed.msi"
    $mongoInstallPath = "C:\Program Files\MongoDB\Server\3.2"
}

$downloadDirectory = "$scriptDir\$mongoInstallFilename"

# download mongo
Start-BitsTransfer -Source $mongoDownloadUrl -Destination $downloadDirectory


# create mongo data directory
md $mongoDataPath


# install mongo
$msiInstallerPath = "$scriptDir\$mongoInstallFilename"

Write-Host "Installing Mongo, please wait..."

Start-Process -FilePath msiexec -ArgumentList /i, $msiInstallerPath, /quiet -Wait

Write-Host "Success! Mongo installed."

# run mongo
$mongoServerPath = "$mongoInstallPath\bin\mongod.exe"

Write-Host "Running Mongo from $mongoServerPath"

$mongoArgs = "--storageEngine=mmapv1 --dbpath $mongoDataPath"
Start-Process $mongoServerPath -ArgumentList $mongoArgs


# import collections
$mainDbName = "StreetSupport"
$testDbName = "StreetSupport_Test"
$collectionsFolderName = "collections"
$collectionsFolderPath = "$scriptDir\$collectionsFolderName"

$mongoRestorePath = "$mongoInstallPath\bin\mongorestore.exe"

$mongoRestoreMainDbArgs = "-d $mainDbName $collectionsFolderPath"
$mongoRestoreTestDbArgs = "-d $testDbName $collectionsFolderPath"

Start-Process $mongoRestorePath -ArgumentList $mongoRestoreMainDbArgs
Start-Process $mongoRestorePath -ArgumentList $mongoRestoreTestDbArgs

Write-Host "Database restore complete!"