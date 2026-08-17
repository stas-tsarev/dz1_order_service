# scripts/install-buf.ps1
param(
    [string]$BinDir = "D:/Gosha/dz1_order_service/bin",
    [string]$BufVersion = "1.72.0"
)

$bufExe = Join-Path $BinDir "buf.exe"

# Check if buf is already installed
if (Test-Path $bufExe) {
    Write-Host "Buf already installed: $bufExe"
    exit 0
}

Write-Host "Installing Buf version $BufVersion..."

# Create necessary directories
New-Item -ItemType Directory -Force -Path $BinDir | Out-Null
New-Item -ItemType Directory -Force -Path "tmp-buf" | Out-Null

try {
    # Download archive
    $url = "https://github.com/bufbuild/buf/releases/download/v$BufVersion/buf-Windows-x86_64.zip"
    $zipPath = "tmp-buf\buf.zip"

    Write-Host "Downloading Buf from $url..."
    Invoke-WebRequest -Uri $url -OutFile $zipPath

    # Extract
    Write-Host "Extracting archive..."
    Expand-Archive -Path $zipPath -DestinationPath "tmp-buf" -Force

    # Move executable
    $sourceExe = "tmp-buf\buf\bin\buf.exe"
    if (Test-Path $sourceExe) {
        Move-Item -Path $sourceExe -Destination $bufExe -Force
        Write-Host "Buf installed successfully: $bufExe"
    } else {
        throw "buf.exe not found in extracted archive"
    }
} catch {
    Write-Error "Error installing Buf: $_"
    exit 1
} finally {
    # Clean up temporary files
    if (Test-Path "tmp-buf") {
        Remove-Item -Recurse -Force "tmp-buf"
        Write-Host "Temporary files cleaned up"
    }
}

# Verify installation
if (Test-Path $bufExe) {
    Write-Host "Buf installed successfully!"
    & $bufExe --version
} else {
    Write-Error "Buf not found after installation"
    exit 1
}