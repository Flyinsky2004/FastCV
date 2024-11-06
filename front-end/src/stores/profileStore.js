import {defineStore} from "pinia";

const useProfileStore = defineStore('profile', {
    state:() => ({
        currentProfile: {},
        isSelected: false
    }),
    actions: {
        setProfile(profile){
            this.currentProfile = profile
            this.isSelected = true
            localStorage.setItem('lastProfile', JSON.stringify(profile))
        },
        setSelected(status){
            this.isSelected = status
        }
    },
    getters:{
        getCurrentProfile:(state) => state.currentProfile
    }
})

export {useProfileStore}