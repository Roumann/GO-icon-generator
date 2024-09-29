# GO icon generator

Icong generator written in go to automatically generate notification icons, app icons for react-native init project.

## Features

### Notification icons (Android)

Generates mdpi, hdpi, xhdpi, xxhdpi and xxxhdpi version of your provided icon to use as default notification icon while using FCM notifications.

### App icons (Android)

Generates mdpi, hdpi, xhdpi, xxhdpi and xxxhdpi and Adaptive icon for your provided icon to use as default app icon.

### Requirements

1. Icon (.png or .jpg)
2. Transparent background
3. Atleast 256x256px

### Usage

App icons
rnGen app <icon path> <background color> -p <0.2-2> [Optional: padding default 0.75]

Notification icons
rnGen notif <icon path> -p <0.2-2> [Optional: padding default 0.75]
