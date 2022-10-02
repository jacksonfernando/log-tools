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
## Screenshot
* json output
![json_output_image](https://github.com/jacksonfernando/log-tools/blob/main/pictures/converted_json.PNG)
* text output
![json_output_image](https://github.com/jacksonfernando/log-tools/blob/main/pictures/converted_text.PNG)


