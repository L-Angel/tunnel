import axios from "axios";
import {host} from "../config";

export function fetch(path, method, data) {
    return new Promise(function (resolve, reject) {
        axios({
            headers: {
                "Content-Type": "application/json"
            },
            method: method,
            url: host() + "" + path,
            data: data
        }).then(resp => {
            resolve(resp)
        }).catch(err => {
            reject(err)
        }).finally(f => {

        })
    })

}
