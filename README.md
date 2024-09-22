# GO icon generator

Icong generator written in go to automatically generate notification icons, app icons for react-native init project.

## Features

Notification icons (Android)

Generates mdpi, hdpi, xhdpi, xxhdpi and xxxhdpi version of your provided icon to use as default notification icon while using FCM notifications.

### Requirements

1. Monochrome Icon (white works nice)
2. Transparent background
3. Atleast 256x256px

### Usage 

--TODO-- go run main.go < path-to-icon > [OPTIONAL]< padding - between 0.5 - 1.5 - default - 0.75 >

Program will generate all sizes and save them to the android/app/src/main/res/drawable-< icon-density > folders. If they dont exist, it will create them for you.
