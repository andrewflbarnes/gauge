#!/usr/bin/env node

"use strict"

const install = require("./install"),
    path = require("path"),
    unzip = require('unzipper'),
    https = require('https'),
    ProxyAgent = require('proxy-agent'),
    packageJsonPath = path.join(__dirname, "..", "package.json"),
    binPath = "./bin";

const downloadFollowingRedirect = function(url, resolve, reject) {
    const opts = {
        headers: { 'accept-encoding': 'gzip,deflate' },
        agent: new ProxyAgent(),
    }
    https.get(url, opts, res => {
        if (res.statusCode >= 300 && res.statusCode < 400) {
            downloadFollowingRedirect(res.headers.location, reject, resolve);
        } else if (res.statusCode > 400) {
            console.error(`Unable to download '${url}' : ${res.statusCode}-'${res.statusMessage}'`);
        } else {
            res.pipe(unzip.Extract({ path: path.normalize(binPath) })).on('error', reject).on('end', resolve);
        }
    });
};

const downloadAndExtract = function(version) {
    console.log(`Fetching download url for Gauge version ${version}`);
    const url = install.getBinaryUrl(version);
    console.log(`Downloading ${url} to ${binPath}`);
    return new Promise((resolve, reject) => {
        try {
            downloadFollowingRedirect(url, resolve, reject);
        } catch (error) {
            reject(error);
        }
    })
};

install.getVersion(packageJsonPath)
    .then((v) => downloadAndExtract(v.split('-')[0]))
    .catch((e) => console.error(e));