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

db.createUser({
  user: "groupsTestUser",
  pwd: "Password456",
  roles: [
    {
      role: "readWrite",
      db: "groups_test",
    },
  ],
});
