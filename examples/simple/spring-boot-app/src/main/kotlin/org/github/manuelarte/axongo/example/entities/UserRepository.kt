package org.github.manuelarte.axongo.example.entities

import org.springframework.data.repository.CrudRepository
import org.springframework.data.repository.PagingAndSortingRepository
import org.springframework.stereotype.Repository

@Repository
interface UserRepository :
    CrudRepository<User, Int>,
    PagingAndSortingRepository<User, Int>
