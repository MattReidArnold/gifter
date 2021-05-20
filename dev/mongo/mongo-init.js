db.createUser({
  user: "groupsUser",
  pwd: "Password123",
  roles: [
    {
      role: "readWrite",
      db: "groups",
    },
  ],
});
