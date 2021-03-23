import { IAppOption } from "../../appoption";
import { TripService } from "../../service/trip";
import { routing } from "../../utils/routing";

const shareLocationKey = "share_location"
Page({
    data:{
        shareLocation:false,
        avatarURL :"",
    },
    async onLoad(opt:Record<'car_id',string>){
        const o: routing.LockOpts = opt
        console.log('unlocking car ',o.car_id);
        const userInfo = await getApp<IAppOption>().globalData.userInfo
        this.setData({
            avatarURL:userInfo.avatarUrl,
            shareLocation:wx.getStorageSync(shareLocationKey)||true,
        })
    },
    onGetUserInfo(e:any){
        const userInfo:WechatMiniprogram.UserInfo = e.detail.userInfo
        if (userInfo){
            getApp<IAppOption>().resolveUserInfo(userInfo)
        }
    },
    onSHareLocation(e:any){
        const shareLocation:boolean = e.detail.value
        wx.setStorageSync(shareLocationKey,shareLocation)
    },
    onUnlockTap(){
        wx.getLocation({
            type:"gcj02",
            success:loc=>{
                console.log('staring a trip',{
                    location:{
                        latitude:loc.longitude,
                        longitude:loc.longitude
                    },
                    //TODO：需要双向绑定
                    avatarURL:this.data.shareLocation?this.data.avatarURL:'',
                })

                TripService.CreateTrip({
                    start:'abc',
                })
                return
                const tripID = 'trip456'

                wx.showLoading({
                    title:"开锁中",
                    mask:true,
                })
                setTimeout(()=>{
                    wx.redirectTo({
                        // url:`/pages/driving/driving?trip_id=${tripID}`,
                        url:routing.drving({
                            trip_id:tripID
                        }),
                        complete:()=>{
                            wx.hideLoading()
                        }
                    })
                   
                },2000)
            },
            fail:()=>{
                wx.showToast({
                    icon:'none',
                    title:'请前往设置页面授权位置信息'
                })
            }
        }) 
    },
})