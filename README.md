# LOG TOOLS
Tools for converting log file to json or plaintext
## Tools usage
* Execute log-tools binary using `./log-tools` command. Example can be seen below.
* usage : `./log-tools [FILE] [OPTIONS]...`
### OPTIONS
```bash
-t flag for target file 
-o flag for output destination
```
## Example Usage
```bash
./log-tools /var/log/dpkg.log
./log-tools /var/log/dplg.log -t json 
./log-tools /var/log/dpkg.log -t text 
./log-tools /var/log/dpkg.log -t json -o ~/dpkg.json
./log-tools /var/log/dpkg.log -t text -o ~/dpkg.text
```
## Notes
* when converting log files to json it only can read to int32(2.147.483.647 lines). Feel free to increase the size in the `main.go` line 62 and 63
![image](https://user-images.githubusercontent.com/33139248/193443133-dea2ec16-e6f2-4fac-a7e5-3699f6dbff45.png)
## Screenshot
* json output
![json_output_image](https://github.com/jacksonfernando/log-tools/blob/main/pictures/converted_json.PNG)
* text output
![json_output_image](https://github.com/jacksonfernando/log-tools/blob/main/pictures/converted_text.PNG)


