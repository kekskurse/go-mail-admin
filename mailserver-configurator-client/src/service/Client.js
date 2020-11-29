import Api from './Api'

export default {
    ping() {
        return Api().get("/ping")
    },
    getStatus() {
        return Api().get("/status")
    },
    getDomains() {
        return Api().get("/api/v1/domain")
    },
    addDomain(domainname) {
      return Api().post("/api/v1/domain", {"domain": domainname})
    },
    removeDomain(domainname) {
        return Api().delete("/api/v1/domain", {"data": {"domain": domainname}})
    },
    getAlias() {
        return Api().get("/api/v1/alias");
    },
    saveAlias(data) {
        return Api().put("/api/v1/alias", data)
    },
    createAlias(data) {
        return Api().post("/api/v1/alias", data)
    },
    removeAlias(id) {
        return Api().delete("/api/v1/alias", {"data": {"id": id}})
    },
    getAccounts() {
        return Api().get("/api/v1/account")
    },
    saveAccount(data) {
        data["quota"] = parseInt(data["quota"])
        return Api().put("/api/v1/account", data)
    },
    createAccount(data) {
        data["quota"] = parseInt(data["quota"])
        return Api().post("/api/v1/account", data)
    },
    deleteAccount(id) {
        return Api().delete("/api/v1/account", {"data": {"id": id}})
    },
    getTLSPolicys() {
        return Api().get("/api/v1/tlspolicy")
    },
    createTLSPolicy(data) {
        return Api().post("/api/v1/tlspolicy", data)
    },
    saveTLSPolicy(data) {
        data["id"] = parseInt(data["id"])
        return Api().put("/api/v1/tlspolicy", data)
    },
    deleteTLSPolicy(id) {
        return Api().delete("/api/v1/tlspolicy", {"data": {"id": id}})
    },
    changePassword(id, newpassword) {
        return Api().put("/api/v1/account/password ", {"id": id, "password": newpassword})
    },
    featureToggles() {
        return Api().get("/public/v1/features")
    },
    login(username, password) {
        return Api().post("/public/v1/login/username", {"username": username, "password": password})
    }
}
