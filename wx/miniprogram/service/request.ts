import camelcaseKeys = require("camelcase-keys")
import { auth } from "./proto_gen/auth/auth_pb"

export namespace Coolcar{
    const serverAddr = 'http://localhost:8080'
    const AUTH_ERR = 'AUTH_ERR'
    const authData = {
        token:'',
        expirMs:0,
    }
    interface RequestOption<REQ,RES>{
        method:'GET'|'PUT'|'POST'|'DELETE'
        path:string
        data:REQ
        respMarshller:(r:object)=>RES
    }

    export interface AuthOption{
        attachAuthHeader:boolean
        retryOnAuthError:boolean
    }

    export async function sendRequestWithAuthRetry<REQ,RES>(o:RequestOption<REQ,RES>,a?:AuthOption):Promise<RES> {
        const authOpt = a||{
            attachAuthHeader:true,
            retryOnAuthError:true,
        }
        try{
            await Login()
            return   sendRequest(o,authOpt)
        }catch(err){
            if (err===AUTH_ERR&&authOpt.retryOnAuthError){
                authData.token = ''
                authData.expirMs=0
                return sendRequestWithAuthRetry(o,{
                    attachAuthHeader:authOpt.attachAuthHeader,
                    retryOnAuthError:false
                })
            }else{
                throw err
            }
        }
    }

    export async function Login() {

            if (authData.token&&authData.expirMs >=Date.now()){
                    return
            }

            const wxResp = await wxLogin()
            const reqTimeMs = Date.now()
            const resp = await sendRequest<auth.v1.ILoginRequest,auth.v1.ILoginResponse>({
                method:'POST',
                path:'/v1/auth/login',
                data:{
                    code:wxResp.code,
                },
                respMarshller:auth.v1.LoginResponse.fromObject
            },{
                attachAuthHeader:false,
                retryOnAuthError:false
            })
            authData.token = resp.accessToken!
            authData.expirMs = reqTimeMs + resp.expiresIn! * 1000
    }


    function sendRequest<REQ,RES>(o:RequestOption<REQ,RES>,a:AuthOption):Promise<RES>{
        return new Promise((resolve,reject)=>{
            const header:Record<string,any>={}
            if(a.attachAuthHeader){
                if(authData.token&&authData.expirMs>=Date.now()){
                    header.authorization ='Bearer '+ authData.token
                }else{
                    reject(AUTH_ERR)
                    return
                }
            }
            wx.request({
                url:serverAddr+o.path,
                method:o.method,
                data:o.data,
                success:res=>{
                    if(res.statusCode===401){
                        reject(AUTH_ERR)
                    }else if(res.statusCode>=400){
                        reject(res)
                    }else{
                        resolve( o.respMarshller(
                            camelcaseKeys(res.data as object,{
                                deep:true
                            })))
                    }
                },
                header,
                fail:reject
            })
        })
    }

    function wxLogin():Promise<WechatMiniprogram.LoginSuccessCallbackResult>{
        return new Promise((resolve,reject)=>{
            wx.login({
                success:resolve,
                fail:reject,
            })
        })
    }

}