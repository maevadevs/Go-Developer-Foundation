# About These Notebooks

These are Jupyter Notebook versions of the contents from the same materials

## To Run Notebook In Docker

These notebooks depends on Docker and [gophernotes](https://github.com/gopherdata/gophernotes)

- Open Terminal
- cd to the `Notebook` folder
- Run command the follwing command

```ps1
docker run -it -p 8888:8888 -v C:\Users\maeva\Desktop\Learn-Series\Learn-Go\Udemy-Course\Notebooks:/notebooks gopherdata/gophernotes:latest-ds
```