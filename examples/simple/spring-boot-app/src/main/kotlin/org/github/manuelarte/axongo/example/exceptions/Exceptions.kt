@file:Suppress("ktlint:standard:filename")

package org.github.manuelarte.axongo.example.exceptions

import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.ControllerAdvice
import org.springframework.web.bind.annotation.ExceptionHandler
import kotlin.reflect.KClass

class IdNotFoundException(
    clazz: KClass<out Any>,
    id: Int,
) : RuntimeException("${clazz.simpleName} with id $id not found")

data class ErrorResponse(
    val status: HttpStatus,
    val message: String?,
)

@ControllerAdvice
class ExceptionHandlerController {
    @ExceptionHandler(IdNotFoundException::class)
    fun notFoundException(e: IdNotFoundException): ResponseEntity<ErrorResponse> =
        ResponseEntity.status(HttpStatus.NOT_FOUND).body(ErrorResponse(HttpStatus.NOT_FOUND, e.message))
}
