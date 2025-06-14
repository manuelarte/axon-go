package org.github.manuelarte.axongo.example.api

import org.github.manuelarte.axongo.example.entities.User


data class UserRead(
    val id: Int,
    val name: String,
    val surname: String,
) {
    companion object Entity {
        fun from(user: User?): UserRead? =
            user?.let {
                UserRead(it.id!!, it.name, it.surname)
            }
    }
}
