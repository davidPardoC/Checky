import AlertComponent from "@/components/Alert/Alert";
import UsersTable from "@/components/UsersTable/UsersTable";
import RolesService from "@/services/roles.services";
import UserServices from "@/services/users.services";
import useUsersStore from "@/store/users.store";
import { Role } from "@/types/roles";
import { User } from "@/types/user";
import { getServerSideAxiosInstance } from "@/utils/server/axios";
import Head from "next/head";

type UsersPageProps = {
  roles: Role[];
  users: User[];
};

export default function Page({ roles, users }: UsersPageProps) {
  const setRoles = useUsersStore((state) => state.setAllroles);
  const setUsers = useUsersStore((state) => state.setAllUsers);
  const error = useUsersStore((state) => state.error);

  setRoles(roles);
  setUsers(users);

  return (
    <>
      {error.show && <AlertComponent message={error.message} />}
      <Head>
        <title>Users</title>
      </Head>
      <UsersTable />
    </>
  );
}

export async function getServerSideProps() {
  const axiosInstance = getServerSideAxiosInstance();
  const roles = await RolesService.getRoles(axiosInstance);
  const users = await UserServices.getAllUsers(axiosInstance);
  return {
    props: { roles, users },
  };
}
