function auth_page() {
  return {
    username: "",
    password: "",
    loading: false,
    action: "login",
    async authenticate() {
      alert("werking")
    }
  }
}
