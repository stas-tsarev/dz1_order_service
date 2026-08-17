# scripts/gen-ogen.ps1
$openAPIFiles = "D:/Gosha/dz1_order_service/shared/api/bundles"
$yq = "D:/Gosha/dz1_order_service/bin/yq.exe"
$ogen = "D:/Gosha/dz1_order_service/bin/ogen.exe"

$files = Get-ChildItem -Path $openAPIFiles -Include *.yaml, *.yml -Recurse

foreach ($file in $files) {
    $content = Get-Content $file.FullName -Raw
    if ($content -match 'x-ogen:') {
        Write-Host "🚀 Generating from: $($file.FullName)"

        $target = & $yq e '."x-ogen".target' $file.FullName
        $package = & $yq e '."x-ogen".package' $file.FullName

        Write-Host "📁 Target: $target"
        Write-Host "📦 Package: $package"

        & $ogen --target $target --package $package --clean $file.FullName

        if ($LASTEXITCODE -ne 0) {
            Write-Error "Failed to generate from $($file.FullName)"
            exit 1
        }
    }
}