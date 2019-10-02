# go-package-plantuml


### Environmental configuration

1.Install go environment and configure environment variables
````ftl>
export GOROOT=/opt/golang/go
export PATH=$GOROOT/bin:$PATH
export GOPATH=/opt/gopath
````
2.Install JDK8 and configure environment variables
````ftl>
export JAVA_HOME=/opt/jdk/jdk1.8.0_161
export PATH=$JAVA_HOME/bin:$PATH
export CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar
````
in /etc/profile --> Add the corresponding environment variable to the file and refresh with source /etc/profile

### Install required software 
````ftl>
yum install graphviz
yum install git
yum install wget
````

### Download and compile the project
````
go get github.com/ahilbig/go-package-plantuml
````

### Compiling and downloading dependencies
First run automatically downloads dependencies, please wait patiently
````
cd /opt
cp $GOPATH/src/github.com/maobuji/go-package-plantuml/goplantuml . -rf
cd goplantuml
chmod 775 *.sh
sh install.sh
````


# Run directly using commands
Direct operation can set more parameters, --codedir must be entered, other parameters are optional
````
./go-package-plantuml --codedir /appdev/gopath/src/github.com/contiv/netplugin \
--gopath /appdev/gopath \
--outputfile  /tmp/result.txt
--ignoredir /appdev/gopath/src/github.com/contiv/netplugin/vendor
````
Parameter Description<br>
--codedir The code directory to analyze<br>
--gopath GOPATH Environment variable directory<br>
--outputfile Analysis results are saved to this file<br>
--ignoredir No need for code analysis directory（Can not set）<br>
--ignorefile, -if No need for  code analysis file
--file, -f Include file - all other will be ignored

The output text of the previous step, convert to svg file
````
java -jar plantuml.jar /tmp/result.txt -tsvg
````

gouml Samples in the script, can be directly run by sh gouml.sh 
