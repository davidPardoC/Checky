import { Role } from "@/types/roles";
import { User } from "@/types/user";
import { create } from "zustand";

type UserPageState = {
  roles: Role[];
  users: User[];
  setAllroles: (roles: Role[]) => void;
  setAllUsers: (roles: User[]) => void;
};

const useUsersStore = create<UserPageState>((set) => ({
  roles: [],
  users: [],
  setAllroles: (roles: Role[]) => set((state) => ({ ...state, roles })),
  setAllUsers: (users: User[]) => set((state) => ({ ...state, users })),
}));

export default useUsersStore;
