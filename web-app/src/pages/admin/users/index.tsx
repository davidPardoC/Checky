import UsersTable from "@/components/UsersTable/UsersTable";
import RolesService from "@/services/roles.services";
import useUsersStore from "@/store/users.store";
import { Role } from "@/types/roles";
import { getServerSideAxiosInstance } from "@/utils/server/axios";
import Head from "next/head";

type UsersPageProps = {
  roles: Role[];
};

export default function Page({ roles }: UsersPageProps) {
  const setRoles = useUsersStore((state) => state.setAllroles);
  setRoles(roles);

  return (
    <>
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
  return {
    props: { roles },
  };
}
