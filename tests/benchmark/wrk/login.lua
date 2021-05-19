local path  = {"/api/user/captcha","/api/user/login","/api/user/info"}
local method = {"GET","POST","GET"}
local hds = {}
local bds = ""
local cur = 0
local bb = {["password"]="666666"}
local students = {
"009vtynx",
"00dema0",
"00l6pcje",
"00rj0eac",
"00uyjti7",
"00xt6gr8",
"012nscu7",
"01310d3",
"01kb407w",
"01sg0d41",
"01xpaaa",
"01yjd2pi",
"01yxtpgd",
"01z48b2d",
"020ar9b",
"026rt80q"}

local s = 0
local cj = require("cjson")

request = function()
   return wrk.format(method[cur+1], path[cur+1],hds,bds)
end

response = function(status, headers, body)
    if status == 200 then
        if cur == 0 then
            bb["account"] = students[s+1]
            s = (s + 1)%16
            local cb = cj.decode(body)
            bb["captchaId"] = cb["data"]["captchaId"]
            bb["picPath"] = cb["data"]["d"]   
            bds = cj.encode(bb)
        elseif cur == 1 then
            local cb = cj.decode(body)
            local token = cb["data"]["token"]
            hds["Authorization"] = token
            bds = nil
        else
            hds = {}
        end
        cur = (cur + 1)%3
    end
end