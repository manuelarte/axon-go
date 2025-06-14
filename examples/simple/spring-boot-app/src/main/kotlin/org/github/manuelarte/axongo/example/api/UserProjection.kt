package org.github.manuelarte.axongo.example.api

import org.axonframework.queryhandling.QueryHandler
import org.github.manuelarte.axongo.example.entities.UserRepository
import org.springframework.context.annotation.Profile
import org.springframework.data.repository.findByIdOrNull
import org.springframework.stereotype.Service

@Profile("query")
@Service
@Suppress("unused")
class UserProjection(
        private val userRepository: UserRepository,
) {
    @QueryHandler
    fun handle(query: GetUserByIDQuery): UserRead? {
        val user = this.userRepository.findByIdOrNull(query.id)
        return UserRead.from(user)
    }
}
