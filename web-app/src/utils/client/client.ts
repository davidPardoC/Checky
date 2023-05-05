const isClient = () => {
    return typeof window !== "undefined"
}

const clientUtils =  {isClient}

export default clientUtils