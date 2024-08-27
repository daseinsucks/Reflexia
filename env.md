# env

## Summary

This code package provides a way to load environment variables from a ".env" file and access them as needed. It uses the "godotenv" library to read the environment variables and store them in a map. The package also includes functions to load specific data, such as admin data, a Telegram token, and a local password.

The package defines a constant "envLoc" that specifies the location of the ".env" file. It also defines a type "AdminData" to represent admin data, which includes an ID and a GPT key. The "Load()" function reads the environment variables from the ".env" file and stores them in the "env" map.

The "LoadAdminData()" function iterates through the environment variables and parses the values for admin data, such as "ADMIN_ID", "MINTY_ID", "OK_ID", and "MURS_ID". It then creates a map of admin data, where each key is the admin ID and the value is an "AdminData" struct containing the ID and GPT key.

The package also includes functions to load a Telegram token, a local password, and an AI endpoint. The "LoadTGToken()" function returns the Telegram token from the environment variable "TG_KEY". The "LoadLocalPD()" function returns the local password from the environment variable "LOCALHOST_PWD". The "LoadLocalAI_Endpoint()" function returns the AI endpoint from the environment variable "AI_ENDPOINT".

Finally, the "GetAdminToken()" function returns the admin token from the environment variable "ADMIN_KEY".



