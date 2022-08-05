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
  methods: {
    created: function () {
      // Simple GET request using fetch
      fetch("http://localhost:8080/")
        .then((response) => response.json())
        .then((data) => (this.emailList = data.map((email) => email.address)))
        .catch((err) => (this.error = err));
    },
    addToList() {
      if (this.email === "") {
        this.error = "Invalid input";
        return;
      }

      this.emailList.unshift(this.email);

      this.email = "";
      this.error = "";
    },
    remove() {
      if (this.email === "") return;
      let idx = this.emailList.indexOf(this.email);
      console.log(this.email);
      if (idx !== -1) this.emailList.splice(idx, 1);
      this.email = "";
    },
  },
};
</script>

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
      <!-- <li v-bind:key="i" v-bind:title="email" v-for="(email, i) in emailList"> -->
      <span>{{ emailList.toString() }}</span>
      <!-- </li> -->
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
