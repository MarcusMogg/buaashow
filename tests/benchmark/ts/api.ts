'use strict';
import * as fetch from "node-fetch"
import * as fs from "fs/promises"
const FormData = require('form-data');

const baseUrl = "http://10.251.253.71";

function randomNum(minNum, maxNum) {
    return Math.floor(Math.random() * (maxNum - minNum + 1)) + minNum;
}

export function randomString(n) {
    let str = 'abcdefghijklmnopqrstuvwxyz9876543210';
    let tmp = '',
        i = 0,
        l = str.length;
    for (i = 0; i < n; i++) {
        tmp += str.charAt(randomNum(0, l));
    }
    return tmp;
}

export class User {
    constructor(public account: string,
        public password: string) {
    }
}

export const captcha = () => {
    const requestOptions = {
        method: 'GET',
        redirect: 'follow'
    };

    return fetch(baseUrl + "/api/user/captcha", requestOptions)
        .then(response => response.json())
        .then(result => result.data)
        .catch(error => console.log('error', error));
}

export const login = (account: string, password: string, captchaId: string, picPath: string) => {
    let myHeaders = new fetch.Headers();
    myHeaders.append("Content-Type", "application/json");

    const raw = JSON.stringify({
        "account": `${account}`,
        "password": `${password}`,
        "captchaId": `${captchaId}`,
        "picPath": `${picPath}`
    });

    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };

    return fetch(baseUrl + "/api/user/login", requestOptions)
        .then(response => response.json())
        .then(result => result.data.token)
        .catch(error => console.log('error', error));
}

export const newTercher = async (token: string) => {
    let myHeaders = new fetch.Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Authorization", token);
    let t = new User("teacher_test_" + randomString(5), randomString(randomNum(6, 12)))
    const raw = JSON.stringify({
        "account": t.account,
        "password": t.password
    });

    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };
    await fetch(baseUrl + "/api/user/teacher", requestOptions)
        .catch(error => console.log('error', error));
    return t
}

export const newTerm = async (token: string,
    tname: string,
    tbegin: string,
    tend: string) => {
    let myHeaders = new fetch.Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Authorization", token);
    const raw = JSON.stringify({
        "tbegin": tbegin,
        "tend": tend,
        "tname": tname
    });

    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };

    return fetch(baseUrl + "/api/terms", requestOptions)
        .then(response => response.json())
        .then(res => res.data.tid)
        .catch(error => console.log('error', error));
}

export const newCN = async (token: string, cname: string) => {
    let myHeaders = new fetch.Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Authorization", token);
    const raw = JSON.stringify({
        "name": cname
    });

    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };

    await fetch(baseUrl + "/api/coursename", requestOptions)
        .catch(error => console.log('error', error));
}

export const newClass = (token: string,
    nameid: number,
    termid: number) => {
    let myHeaders = new fetch.Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Authorization", token);
    //console.log(termid);
    const raw = JSON.stringify({
        "name_id": nameid,
        "info": randomString(10),
        "tid": termid
    });

    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };

    return fetch(baseUrl + "/api/course", requestOptions)
        .then(response => response.json())
        .then(res => { return res.data.cid })
        .catch(error => console.log('error', error));
}

export const newExp = (token: string, cid: Number) => {
    let myHeaders = new fetch.Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Authorization", token);
    const raw = JSON.stringify({
        "info": randomString(10),
        "name": randomString(10),
        "team": false
    });

    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };
    return fetch(baseUrl + `/api/course/${cid}/exp`, requestOptions)
        .then(response => response.json())
        .then(res => res.data.eid)
        .catch(error => console.log('error', error));
}

export const addStudents = async (token: string,
    cid: number,
    students: string[]) => {
    let myHeaders = new fetch.Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Authorization", token);
    const raw = JSON.stringify({
        "accounts": students
    });

    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };

    await fetch(baseUrl + `/api/course/${cid}/students`, requestOptions)
        .catch(error => console.log('error', error));
}

export const postFile = (path: string) => {
    let formdata = new FormData();
    return fs.readFile(path)
        .then(res => formdata.append("file", res, path)).
        then(() => {
            return {
                method: 'POST',
                body: formdata,
                redirect: 'follow'
            }
        })
        .then(x => fetch(baseUrl + "/api/file", x))
        .then(response => response.json())
        .then(result => { return result.data })
        .catch(err => console.error(err))
}

export const postImg = (path: string) => {
    let formdata = new FormData();
    return fs.readFile(path)
        .then(res => formdata.append("file", res, path)).
        then(() => {
            return {
                method: 'POST',
                body: formdata,
                redirect: 'follow'
            }
        })
        .then(x => fetch(baseUrl + "/api/img?width=740&height=400", x))
        .then(response => response.json())
        .then(result => { return result.data })
        .catch(err => console.error(err))
}

export const submit = async (token: string,
    dist: string,
    thumb: string,
    eid: number) => {
    let myHeaders = new fetch.Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Authorization", token);
    const raw = JSON.stringify({
        "info": randomString(5),
        "name": randomString(15),
        "readme": randomString(150),
        "type": 1,
        "src_url": "",
        "dist_url": dist,
        "thumb": thumb
    });
    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };

    await fetch(baseUrl + `/api/exp/${eid}/submit`, requestOptions)
        .catch(error => console.log('error', error));
}

export const rec = async (token: string,
    account: string,
    eid: number) => {
    let myHeaders = new fetch.Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Authorization", token);
    const raw = JSON.stringify({
        "account": account,
        "rec": true
    });

    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };

    await fetch(baseUrl + `/api/exp/${eid}/rec`, requestOptions)
        .catch(error => console.log('error', error));
}