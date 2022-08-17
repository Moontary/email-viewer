<script>
export default {
  data() {
    return {
      error: "",
      email: "",
      emailList: [],
    };
  },
  mounted() {
    this.created();
  },
  // computed props
  computed: {
    list: function () {
      return this.emailList.map((data) => data.address).toString();
    },
  },
  methods: {
    created: function () {
      // Simple GET request using fetch
      const URL = import.meta.env.VITE_URL;
      fetch(URL, { method: "GET" })
        .then((response) => response.json())
        .then((data) => (this.emailList = data))
        .catch((err) => (this.error = err));
    },
    // addToList adds input email to the list
    addToList: function () {
      if (this.email === "") {
        this.error = "Invalid input";
        return;
      }
      fetch(import.meta.env.VITE_URL, {
        method: "POST",
        body: JSON.stringify({ address: this.email }),
      })
        .then((response) => {
          if (response.ok) {
            return response.json();
          }
          return Promise.reject(response.statusText);
        })
        .then((data) => {
          this.emailList.unshift(data);
          this.email = "";
        })
        .catch((err) => (this.error = err));

      this.error = "";
    },
    // remove deletes email from the list
    remove() {
      if (this.email === "") {
        this.error = "Input can't be empty";
        return;
      }
      let email = this.emailList.find((e) => e.address === this.email);
      if (email == null) {
        this.error = "email doesn't exist";
        return;
      }
      fetch(`http://127.0.0.1:9090/${email.id}`, { method: "DELETE" })
        .then((response) => {
          if (response.ok) {
            let idx = this.emailList.findIndex((e) => e.address === this.email);
            console.log(idx);
            if (idx !== -1) this.emailList.splice(idx, 1);
            this.email = "";
            return;
          }
          return Promise.reject(response.statusText);
        })
        .catch((err) => (this.error = err));
    },
  },
};
</script>
<!--template with buttons and forms description-->
<template>
  <div>
    <label>
      <input
        @keyup.enter="addToList"
        type="text"
        v-model="email"
        placeholder="Type the email address"
      />
      <button @click="addToList">Add</button>
      <button @click="remove">Delete</button>
      <div>
        {{ error }}
      </div>
    </label>
    <ul>
      <span>{{ list }}</span>
    </ul>
  </div>
</template>

<style scoped>
ul {
  display: table;
  width: 100%;
  text-align: center;
}

ul > li {
  display: table-cell;
}
</style>
