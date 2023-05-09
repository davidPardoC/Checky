import { Role } from "@/types/roles";
import { User } from "@/types/user";
import { create } from "zustand";

type UserPageState = {
  roles: Role[];
  users: User[];
  error: { show: boolean; message: string };
  setAllroles: (roles: Role[]) => void;
  setAllUsers: (roles: User[]) => void;
  showError: (message: string) => void;
  hideError: () => void;
};

const useUsersStore = create<UserPageState>((set, get) => ({
  error: { show: false, message: "" },
  roles: [],
  users: [],
  setAllroles: (roles: Role[]) => set((state) => ({ ...state, roles })),
  setAllUsers: (users: User[]) => set((state) => ({ ...state, users })),
  showError: (message: string) => {
    set((state) => ({ ...state, error: { show: true, message } }));
    const interval = setTimeout(()=>{
      get().hideError()
      clearInterval(interval)
    }, 1000)
  },
  hideError: () =>
    set((state) => ({ ...state, error: { show: false, message: "" } })),
}));

export default useUsersStore;
