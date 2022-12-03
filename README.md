# gzip2parts

Generates gzipped parts from folder.  

It can be use for uploading big files or folders. When compression running you can access finished parts.  

```
Usage of bin/gzip2parts-{platformname}:  
  -c	Compress Folder 
  -i string 
    	InputFolder 
  -o string 
    	OutputFolder 
  -ps int 
    	PartSize (default 3072000) 
  -x	Exract Folder 
  ```
  
  You can build application with ``` make ```  
  You can run test with ``` make run ```
  
  ![gzip2parts.gif](https://github.com/fatihsoydan/gzip2parts/blob/master/testContent/gzip2parts.gif?raw=true)
  
  You can find specially usage examples of ```compress/gzip``` , ```encoding/gob``` in this project.
  
  
