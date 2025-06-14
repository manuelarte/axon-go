package org.github.manuelarte.axongo.example.controllers

import org.axonframework.extensions.kotlin.query
import org.axonframework.queryhandling.QueryGateway
import org.github.manuelarte.axongo.example.api.GetUserByIDQuery
import org.github.manuelarte.axongo.example.api.UserRead
import org.github.manuelarte.axongo.example.exceptions.IdNotFoundException
import org.springframework.context.annotation.Profile
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController
import java.util.concurrent.CompletableFuture

@RestController
@Profile("api")
@RequestMapping("/users/{id}")
class QueryUserController(
    val queryGateway: QueryGateway,
) {
    @GetMapping("")
    fun getOne(
        @PathVariable id: Int,
    ): CompletableFuture<ResponseEntity<UserRead>> =
        queryGateway
            .query<UserRead, GetUserByIDQuery>(
                GetUserByIDQuery(id),
            ).thenApply {
                if (it == null) {
                    throw IdNotFoundException(UserRead::class, id)
                }
                ResponseEntity.ok(it)
            }
}
