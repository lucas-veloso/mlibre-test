# examen mercadolibre

## Ambiente

Las pruebas se corrieron utilizando la siguiente configuracion de hardware y software

steps:

1-  CentOS7 release 7.3.1611
2-  Mozilla Firefox 52.2.0
3-  Geckodriver v0.18.0
4-  Selenium-webdriver 2.53.4
5-  Ruby 2.0.0
6-  Minitest testing framework

## pasos a seguir para instaslar un ambiente

En una estacion o maquina virtual con CentOS instalado ejecutar primero que nada
```
sudo yum update -y
```
luego chequeamos que tengamos ruby instalado y con la version correcta ejecutando

```
ruby -v 
```
en caso de que no basta con ejecutar

```
sudo yum install ruby ruby-devel rubygems 
```

luego instalar las gemas necesarias

```
sudo gem install selenium-webdriver minitest 
```

solo queda descargar e agregar la ruta del geckodriver 

```
wget https://github.com/mozilla/geckodriver/releases/download/v0.18.0/geckodriver-v0.18.0-arm7hf.tar.gz
tar xf geckodriver-v0.18.0-arm7hf.tar.gz
export PATH=$PATH:ruta/al/geckodriver
```
para correr las pruebas solo queda

```
ruby stest.rb 
```
