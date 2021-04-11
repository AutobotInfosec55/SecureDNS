from tkinter import *
from scapy.all import *
from base64 import b64decode
import tkinter.messagebox
import requests
import uuid
import time
import os

def login():
    global root2
    root2 = Toplevel(root)
    root2.title("Account Login")
    root2.geometry("500x500")
    root2.config(bg="white")

    global username_verification
    global password_verification
    global has_mentor
    Label(root2, text='Please Enter your Mentee Details', bd=5,font=('arial', 12, 'bold'), relief="flat", fg="white",
                   bg="#00C9A7",width=300).pack()
    username_verification = StringVar()
    password_verification = StringVar()
    has_mentor = StringVar()
    Label(root2, text="").pack()
    Label(root2, text="Username ", fg="#00C9A7", font=('arial', 12, 'bold')).pack()
    Entry(root2, textvariable=username_verification).pack()
    Label(root2, text="").pack()

    Label(root2, text="Password ", fg="#00C9A7", font=('arial', 12, 'bold')).pack()
    Entry(root2, textvariable=password_verification, show="*").pack()
    Label(root2, text="").pack()

    Label(root2, text="Mentor Email (leave blank if you're an mentor)", fg="#00C9A7", font=('arial', 12, 'bold')).pack()
    Label(root2, text="").pack()
    Entry(root2, textvariable=has_mentor).pack()
    Label(root2, text="").pack()

    Button(root2, text="Login", bg="#00C9A7", fg='white', relief="flat", font=('arial', 12, 'bold'),command=get_jwt_token).pack()
    Label(root2, text="")

def logged_destroy():
    logged_message.destroy()
    root2.destroy()

def failed_destroy():
    failed_message.destroy()

def logged():
    global logged_message
    logged_message = Toplevel(root2)
    logged_message.title("Welcome")
    logged_message.geometry("500x100")
    Label(logged_message, text="Login Successfully!... Welcome {} ".format(_usernameverification.get()), fg="green", font="bold").pack()
    Label(logged_message, text="").pack()
    Button(logged_message, text="Logout", bg="#00C9A7", fg='white', relief="flat", font=('arial', 12, 'bold'), command=logged_destroy).pack()


def failed():
    global failed_message
    failed_message = Toplevel(root2)
    failed_message.title("Invalid Message")
    failed_message.geometry("500x100")
    Label(failed_message, text="Invalid Username or Password", fg="red", font="bold").pack()
    Label(failed_message, text="").pack()
    Button(failed_message,text="Ok", bg="#00C9A7", fg='white', relief="flat", font=('arial', 12, 'bold'), command=failed_destroy).pack()


def get_jwt_token():
    global jwt
    mac = (':'.join(['{:02x}'.format((uuid.getnode() >> ele) & 0xff)for ele in range(0,8*6,8)][::-1]))
    user_verification = username_verification.get()
    pass_verification = password_verification.get()
    mentor = has_mentor.get()
    if mentor=="":
        mentor = "false"
    data = {"email": user_verification, "password": pass_verification, "has_mentor": mentor,"macaddress":mac}
    print(data)
    response = requests.post('http://101.53.147.32:8080/login',data=data)
    jwt = response.text
    if "Invalid Username or Password" in jwt:
        failed()
    else:
        #this should be a infinte loop
        for i in range(100):
            captureDNS()

def captureDNS():
    global jwt
    packets = sniff(filter='ip host 101.53.147.32',count=20)
    data = ""
    if len(jwt.split('.')[1])%4 != 0:
        tmp = str(jwt.split('.')[1]) + '='*(4-(len(jwt.split('.')[1])%4))
    else:
        tmp= str(jwt.split('.')[1])
    tmp = b64decode(tmp)

    print (tmp)
    for i in packets:
        data = data + "\n" + str(time.time()) + " " + tmp + "  " + i.command()
    sendAnalytics(data)


def sendAnalytics(data):
    global jwt
    headers = { "Authorization": "bearer "+jwt}
    r = requests.post("http://101.53.147.32:8080/log",headers=headers,data=data)
    print('sent log')


def main_display():
    global root
    root = Tk()
    root.config(bg="white")
    root.title("DNS Authentication Panel")
    root.geometry("500x500")
    Label(root,text='Welcome to CPF DNS',  bd=20, font=('arial', 20, 'bold'), relief="flat", fg="white",
                   bg="lightgreen",width=300).pack()
    Label(root,text="").pack()
    Button(root,text='Log In', height="1",width="20", bd=8, font=('arial', 12, 'bold'), relief="flat", fg="white",
                   bg="#00C9A7",command=login).pack()
    
    
main_display()
root.mainloop()
