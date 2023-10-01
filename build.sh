#!/bin/bash
cd client && flutter build macos
mv ./build/macos/Build/Products/Release/client.app ../Notifier.app/Contents/MacOS/client.app
cd ..
cd server && go build
mv ./notifier ../Notifier.app/Contents/MacOS/notifier
cd ..

rm Notifier.app/Contents/MacOS/notifier_error.log
rm Notifier.app/Contents/MacOS/notifier_output.log

hdiutil create -volname "Notifier" -srcfolder "Notifier.app" -ov -format UDZO "Notifier.dmg"
