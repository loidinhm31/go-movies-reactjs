import { Direction, PageType } from "@/types/page";
import { UserType } from "@/types/users";

export const fakeUsers: PageType<UserType> = {
  size: 5,
  page: 0,
  sort: {
    orders: [
      {
        property: "created_at",
        direction: Direction.ASC
      }
    ]
  },
  total_elements: 5,
  total_pages: 1,
  content: [
    {
      id: 1,
      username: "test1",
      email: "test1@example.com",
      first_name: "TestL1",
      last_name: "TestF1",
      is_new: false,
      created_at: "2023-06-02T04:03:13.075557Z",
      role: {
        role_name: "admin",
        role_code: "ADMIN"
      }
    },
    {
      id: 2,
      username: "test2",
      email: "test2@example.com",
      first_name: "TestL2",
      last_name: "TestF2",
      is_new: false,
      created_at: "2023-06-02T04:03:13.075557Z",
      role: {
        role_name: "general",
        role_code: "GENERAL"
      }
    },
    {
      id: 3,
      username: "test3",
      email: "test3@example.com",
      first_name: "TestL3",
      last_name: "TestF3",
      is_new: false,
      created_at: "2023-06-02T04:03:13.075557Z",
      role: {
        role_name: "general",
        role_code: "GENERAL"
      }
    },
    {
      id: 4,
      username: "test4",
      email: "test4@example.com",
      first_name: "TestL4",
      last_name: "TestF3",
      is_new: true,
      created_at: "2023-06-02T04:03:13.075557Z",
      role: {
        role_name: "general",
        role_code: "GENERAL"
      }
    },
    {
      id: 5,
      username: "test5",
      email: "test5@example.com",
      first_name: "TestL5",
      last_name: "TestF5",
      is_new: true,
      created_at: "2023-06-02T04:03:13.075557Z",
      role: {
        role_name: "moderator",
        role_code: "MOD"
      }
    }
  ]
};