rm debug.keystore
keytool -genkey -v -validity 365 -keystore debug.keystore -storepass android -alias androiddebugkey -keypass android -dname "CN=Android Debug,O=Android,C=US"
