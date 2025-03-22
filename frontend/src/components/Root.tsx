import React, { useEffect, useState } from "react"
import { Link } from "react-router-dom";
import Cookies from "js-cookie";

export const Root=()=>{
    const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
    useEffect(() => {
        const token = Cookies.get("token");
        if (token) {
            setIsLoggedIn(() => true);
        } else {
            setIsLoggedIn(() => false);
        }
    }, []);
    return(
        <>
          <div className="w-full h-[90vh] flex flex-col justify-between ">
            <div className="flex flex-row justify-between items-center">
                <h1 className="text-3xl font-semibold">Inventory Management System</h1>
                {isLoggedIn
                    ? <Link to="dashboard"
                            className="rounded-lg px-4 transition-all bg-black py-[0.55rem] text-white hover:bg-[#434343]">Dashboard
                    </Link>
                    : <div className="flex flex-row items-center justify-between md:space-x-6 ">
                        <Link to="signup"
                              className="rounded-lg border-2 border-solid bg-transparent px-4 py-2 transition-all border-black text-black hover:bg-[#dedede]"
                        >Sign Up
                        </Link>
                        <Link to="login"
                              className="rounded-lg px-4 transition-all bg-black py-[0.55rem] text-white hover:bg-[#434343]">Login
                        </Link>
                    </div>
                }
            </div>
        </div>
        </>
    )
}