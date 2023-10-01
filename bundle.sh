#!/bin/bash
hdiutil create -volname "Notifier" -srcfolder "Notifier.app" -ov -format UDZO "Notifier.dmg"
