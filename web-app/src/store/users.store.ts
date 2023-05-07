import { Role } from "@/types/roles";
import { create } from "zustand";

type UserPageState = {
  roles: Role[];
  setAllroles: (roles: Role[]) => void
};

const useUsersStore = create<UserPageState>((set) => ({
  roles: [],
  setAllroles: (roles: Role[]) => set(() => ({ roles })),
}));

export default useUsersStore;
