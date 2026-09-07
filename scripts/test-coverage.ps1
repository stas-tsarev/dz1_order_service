#!/usr/bin/env pwsh

$COVERAGE_DIR = "coverage"

Write-Host "🧪 Запускаем тесты с покрытием..." -ForegroundColor Cyan
Remove-Item -Recurse -Force $COVERAGE_DIR -ErrorAction SilentlyContinue
New-Item -ItemType Directory -Path $COVERAGE_DIR -Force | Out-Null

# Список ваших модулей
$modules = @('order', 'payment', 'inventory')

foreach ($mod in $modules) {
    if (Test-Path "./$mod") {
        Write-Host "📦 Тестируем модуль: $mod" -ForegroundColor Yellow
        go test -cover -coverprofile="$COVERAGE_DIR/$mod.out" -covermode=atomic "./$mod/..."

        if ($LASTEXITCODE -ne 0) {
            Write-Host "❌ Ошибки в модуле $mod" -ForegroundColor Red
            exit $LASTEXITCODE
        }
    }
}

Write-Host ""
Write-Host "📊 Покрытие по каждому модулю:" -ForegroundColor Green
$outFiles = Get-ChildItem "$COVERAGE_DIR/*.out" -ErrorAction SilentlyContinue
foreach ($file in $outFiles) {
    $mod = $file.BaseName
    $line = go tool cover -func=$file | Select-Object -Last 1
    Write-Host " • $mod: $line"
}

Write-Host ""
Write-Host "📦 Общее покрытие:" -ForegroundColor Green
if ($outFiles.Count -gt 0) {
    go tool cover -func="$COVERAGE_DIR/*.out" | Select-Object -Last 1
}

Write-Host ""
Write-Host "📁 Отчёт сохранён в: $COVERAGE_DIR" -ForegroundColor Green