// import { routing } from "../../utils/routing"

import { routing } from "../../utils/routing"

Page({

    redirectURL: '',
    data: {
        licNo:"",
        name:"",
        birthDate:'1990-01-01',
        genderIndex:0,
        genders:['未知','男','女','其他'],
        licImgURL: '',
        state:'UNSUBMITTED' as 'PENDING'|'UNSUBMITTED'|'VERIFIED',
    },
    onLoad(opt:Record<'redirect',string>) {
        const o:routing.RegisterOpts = opt
        if(o.redirect) {
            this.redirectURL = decodeURIComponent(o.redirect)
        }
    }, 
    onUploadLic() {
        wx.chooseImage({
            success: res => {
                if (res.tempFilePaths.length > 0) {
                    this.setData({
                        licImgURL: res.tempFilePaths[0]
                    })
                    //TODO: upload image
                    setTimeout(()=>{
                        this.setData({
                            licNo:"123123123123",
                            name:"张三",
                            birthDate:'1989-02-02',
                            genderIndex:1,
                        })
                    },1000)
                }
            },
        })
    },

    onGenderChange(e: any){
        this.setData({
            genderIndex:e.detail.value
        })
    },

    onBrthDateChange(e: any){
        this.setData({
            birthDate:e.detail.value
        })
    },
    onSubmit(){
        //TODO:submit the form to server
        this.setData({
            state:'PENDING',
        })
        setTimeout(()=>{
            this.onLicVerified()
        },3000)
    },

    onResubmit(){
        //TODO:submit the form to server
        this.setData({
            state:'UNSUBMITTED',
            licImgURL:'',
        })
    },
    onLicVerified() {
        this.setData({
            state: 'VERIFIED',
        })
        if (this.redirectURL) {
            wx.redirectTo({
                url: this.redirectURL,
            })
        }
    }

})