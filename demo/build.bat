@echo off
:: Build Script for Cross-Platform Builds with Go

:menu
cls
echo ===============================
echo Choose the target platform (GOOS) for building:
echo 1. Windows (amd64/windows)
echo 2. Linux (arm/linux)
echo 3. Linux (aarch64/linux)
echo 4. MacOS (amd64/darwin)
echo 5. Android (armv7/android)
echo 6. Android (aarch64/android)
echo ===============================
set /p choice=Enter your choice (1-6): 

:: Validate input
if "%choice%"=="1" goto build_windows
if "%choice%"=="2" goto build_linux_arm
if "%choice%"=="3" goto build_linux_aarch64
if "%choice%"=="4" goto build_macos
if "%choice%"=="5" goto build_android_armv7
if "%choice%"=="6" goto build_android_aarch64

echo Invalid choice, please select 1 - 6.
pause
goto menu

:build_windows
echo Building for Windows (amd64)...
set GOARCH=amd64
set GOOS=windows
goto build

:build_linux_arm
echo Building for Linux (arm)...
set GOARCH=arm
set GOOS=linux
goto build

:build_linux_aarch64
echo Building for Linux (aarch64)...
set GOARCH=arm64
set GOOS=linux
goto build

:build_macos
echo Building for MacOS (amd64)...
set GOARCH=amd64
set GOOS=darwin
goto build

:build_android_armv7
echo Building for Android (armv7)...
set GOARCH=arm
set GOARM=7
set GOOS=android
set CGO_ENABLED=1
set ANDROID_NDK=C:\Users\Gebruiker\AppData\Local\Android\Sdk\ndk\28.0.12674087
set CC=%ANDROID_NDK%/toolchains/llvm/prebuilt/windows-x86_64/bin/armv7a-linux-androideabi21-clang
goto build

:build_android_aarch64
echo Building for Android (aarch64)...
set GOARCH=arm64
set GOOS=android
set CGO_ENABLED=1
set ANDROID_NDK=C:\Users\Gebruiker\AppData\Local\Android\Sdk\ndk\28.0.12674087
set CC=%ANDROID_NDK%/toolchains/llvm/prebuilt/windows-x86_64/bin/aarch64-linux-android21-clang
goto build

:build
:: Execute the Go build command
go build
if errorlevel 1 (
    echo Build failed.
    pause
    goto end
) else (
    echo Build succeeded.
    pause
    goto end
)

:end
